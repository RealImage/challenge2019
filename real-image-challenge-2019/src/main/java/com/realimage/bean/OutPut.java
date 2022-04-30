package com.realimage.bean;

import lombok.Data;

@Data
public class OutPut {
	private String deliveryId;
	private boolean isDeliveryPossible;
	private String partnetId;
	private int costOfDelivery;

}
