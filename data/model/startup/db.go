package startup

import (
	"metaLand/data/model/tag"

	"gorm.io/gorm"
)

// ListStartups  list startups
func ListStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false")
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	// err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Find(startups).Error
	return
}

// StartupNameIsExist check startup's  name is existed
func StartupNameIsExist(db *gorm.DB, name string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("name = ?", name).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// StartupTokenContractIsExist check startup's  token contract is existed
func StartupTokenContractIsExist(db *gorm.DB, tokenContract string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("token_contract_address = ?", tokenContract).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// CreateStartup  create startup
func CreateStartup(db *gorm.DB, startup *Startup) (err error) {
	return db.Create(startup).Error
}

// ListStartupTeamMembers  get startup team members
func ListStartupTeamMembers(db *gorm.DB, startupID uint64, input *ListStartupTeamMemberRequest, output *[]*StartupTeamMember) (total int64, err error) {
	if startupID != 0 {
		db = db.Where("startup_id = ?", startupID)
	}
	if err = db.Table("startup_team_member_rel").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at ASC").Limit(input.Limit).Offset(input.Offset).Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills", "category = ?", tag.ComerSkill).Find(output).Error
	return
}

// CreateStartupTeamMembers  add startup team members
func CreateStartupTeamMembers(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Create(input).Error
}

// GetStartup get startup by startupID
func GetStartup(db *gorm.DB, startupID uint64, output *Startup) (err error) {
	if startupID == 0 {
		return gorm.ErrRecordNotFound
	}
	// err = db.Where("id = ?", startupID).Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").First(output).Error
	err = db.Where("id = ?", startupID).First(output).Error

	if err != nil {
		return
	}
	if output.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetStartupsByComerID get startups by comerID
