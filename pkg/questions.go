package pkg

import "github.com/AlecAivazis/survey/v2"

func Questions(options []string)(string,error) {
	qs := []*survey.Question{
		
			{
				Name: "color",
			Prompt: &survey.Select{
				Message: "Choose a file:",
				Options: options,
				
			},
			Validate: survey.Required,
			},
		}
		var ans string;

	err:=survey.Ask(qs,&ans);
	if err!=nil {
		return ans,err
	}
	return ans,nil
}

	

