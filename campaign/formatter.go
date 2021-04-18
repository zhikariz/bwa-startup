package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) (campaignFormatter CampaignFormatter) {
	campaignFormatter.ID = int(campaign.ID)
	campaignFormatter.UserID = int(campaign.UserID)
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return
}

func FormatCampaigns(campaign []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaign {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
