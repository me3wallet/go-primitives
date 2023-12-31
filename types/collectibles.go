package types

var (
	ExtensionMimetypes = map[string]string{
		"png":  "image/png",
		"jpg":  "image/jpg",
		"jpeg": "image/jpg",
		"gif":  "image/gif",
	}
)

type (
	Collection struct {
		Id                string            `json:"id"`
		Name              string            `json:"name"`
		ImageUrl          string            `json:"image_url"`
		Description       string            `json:"description"`
		ExternalLink      string            `json:"external_link"`
		Total             int               `json:"total"`
		Address           string            `json:"address"`
		Coin              uint              `json:"coin"`
		Type              string            `json:"-"`
		PreviewImageURL   *CollectibleMedia `json:"preview_image_url,omitempty"`
		OriginalSourceURL CollectibleMedia  `json:"original_source_url"`
	}

	CollectionPage []Collection

	CollectibleTransferFee struct {
		Asset  string `json:"asset"`
		Amount string `json:"amount"`
	}

	CollectibleMedia struct {
		Mimetype string `json:"mimetype,omitempty"`
		URL      string `json:"url"`
	}

	Collectible struct {
		ID                string                  `json:"id"`
		CollectionID      string                  `json:"collection_id"`
		TokenID           string                  `json:"token_id"`
		ContractAddress   string                  `json:"contract_address"`
		Category          string                  `json:"category"`
		ImageUrl          string                  `json:"image_url"`
		ExternalLink      string                  `json:"external_link"`
		ProviderLink      string                  `json:"provider_link"`
		Type              string                  `json:"type"`
		Description       string                  `json:"description"`
		Coin              uint                    `json:"coin"`
		Name              string                  `json:"name"`
		Version           string                  `json:"nft_version"`
		TransferFee       *CollectibleTransferFee `json:"transfer_fee,omitempty"`
		PreviewImageURL   *CollectibleMedia       `json:"preview_image_url,omitempty"`
		OriginalSourceURL CollectibleMedia        `json:"original_source_url"`
		Properties        []CollectibleProperty   `json:"properties"`
		About             string                  `json:"about"`
		Balance           string                  `json:"balance"`
	}

	CollectibleProperty struct {
		Key    string  `json:"key"`
		Value  string  `json:"value"`
		Rarity float32 `json:"rarity,omitempty"`
	}

	CollectiblePage []Collectible
)
