package model

type CVWrapperModel struct {
	CV []CVModel `json:"cv"`
}

type CVModel struct {
	Basics 			BasicsModel 		`json:"basics"`
	Work   			[]WorkModel			`json:"work"`
	Volunteer 		[]VolunteerModel 	`json:"volunteer"`
	Education 		[]EducationModel 	`json:"education"`
	Awards 			[]AwardModel 		`json:"awards"`
	Publications 	[]PublicationModel 	`json:"publications"`
	Skills 			[]SkillModel 		`json:"skills"`
	Languages 		[]LanguageModel 	`json:"languages"`
	Interests 		[]InterestModel 	`json:"interests"`
	References 		[]ReferenceModel 	`json:"references"`
	Projects 		[]ProjectModel 		`json:"projects"`
}