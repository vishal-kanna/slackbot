package types

type Proposals struct {
	Proposals []struct {
		ProposalID string `json:"proposal_id"`
		Content    struct {
			Type        string `json:"@type"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Changes     []struct {
				Subspace string `json:"subspace"`
				Key      string `json:"key"`
				Value    string `json:"value"`
			} `json:"changes"`
		} `json:"content,omitempty"`
		Status           string `json:"status"`
		FinalTallyResult struct {
			Yes        string `json:"yes"`
			Abstain    string `json:"abstain"`
			No         string `json:"no"`
			NoWithVeto string `json:"no_with_veto"`
		} `json:"final_tally_result"`
		SubmitTime     string `json:"submit_time"`
		DepositEndTime string `json:"deposit_end_time"`
		TotalDeposit   []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"total_deposit"`
		VotingStartTime string `json:"voting_start_time"`
		VotingEndTime   string `json:"voting_end_time"`
	} `json:"proposals"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

// QueryParams to map the query params of an url
type QueryParams map[string]string

// HTTPOptions of a target
type HTTPOptions struct {
	Endpoint    string
	QueryParams QueryParams
	Body        []byte
	Method      string
}
