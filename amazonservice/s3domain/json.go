package s3domain

type JsonFile struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
	Data    struct {
		Inv struct {
			ID             int    `json:"id"`
			Title          string `json:"title"`
			Created        string `json:"created"`
			Updated        string `json:"updated"`
			Warehouse      string `json:"warehouse"`
			Message        string `json:"message"`
			Type           string `json:"type"`
			Status         string `json:"status"`
			Attachment     string `json:"attachment"`
			Detail         string `json:"detail"`
			Seller         string `json:"seller"`
			SellerName     string `json:"seller_name"`
			SenderName     string `json:"sender_name"`
			SenderPhone    string `json:"sender_phone"`
			SenderPhoto    string `json:"sender_photo"`
			ReceiverName   string `json:"receiver_name"`
			ReceiverAddr   string `json:"receiver_addr"`
			ReceiverPhone  string `json:"receiver_phone"`
			ReceiverLatlon string `json:"receiver_latlon"`
			Info           string `json:"info"`
			PickupType     string `json:"pickup_type"`
			PickupStatus   string `json:"pickup_status"`
			ReturnStatus   string `json:"return_status"`
			Logistic       struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Type        string `json:"type"`
				Active      string `json:"active"`
				Cashless    string `json:"cashless"`
				BookingType string `json:"booking_type"`
				Company     string `json:"company"`
			} `json:"logistic"`
			TransactionVal int    `json:"transaction_val"`
			Fee            string `json:"fee"`
			FeeWarehouse   int    `json:"fee_warehouse"`
			BookingCode    string `json:"booking_code"`
			StorageType    string `json:"storage_type"`
			ShowLogo       string `json:"show_logo"`
			Packaging      []struct {
				PackagingID         string `json:"packaging_id"`
				PackagingName       string `json:"packaging_name"`
				PackagingUnit       string `json:"packaging_unit"`
				PackagingPrice      string `json:"packaging_price"`
				PackagingImage      string `json:"packaging_image"`
				PackagingImageThumb string `json:"packaging_image_thumb"`
				Qty                 string `json:"qty"`
				Total               string `json:"total"`
			} `json:"packaging"`
			PackagingVal  string `json:"packaging_val"`
			PackagingSubs string `json:"packaging_subs"`
			Items         []struct {
				ProductID    int         `json:"product_id"`
				SkuID        int         `json:"sku_id"`
				SkuLabel     string      `json:"sku_label"`
				Name         string      `json:"name"`
				Qty          int         `json:"qty"`
				Price        int         `json:"price"`
				Picture      string      `json:"picture"`
				PictureThumb string      `json:"picture_thumb"`
				Qc           string      `json:"qc"`
				QcStatus     string      `json:"qc_status"`
				QcPass       string      `json:"qc_pass"`
				BinPicked    interface{} `json:"bin_picked"`
				ExpiredFlag  string      `json:"expired_flag"`
			} `json:"items"`
			Exchanged          interface{} `json:"exchanged"`
			PerformanceTime    int         `json:"performance_time"`
			PerformanceCouting bool        `json:"performance_couting"`
			IsLock             string      `json:"is_lock"`
			Marketplace        string      `json:"marketplace"`
			MarketplaceName    string      `json:"marketplace_name"`
			Platform           string      `json:"platform"`
			PickingTs          string      `json:"picking_ts"`
			PackingTs          string      `json:"packing_ts"`
			ShippingTs         string      `json:"shipping_ts"`
			DoneTs             string      `json:"done_ts"`
			IntervalAwb        int         `json:"interval_awb"`
		} `json:"inv"`
	} `json:"data"`
}
