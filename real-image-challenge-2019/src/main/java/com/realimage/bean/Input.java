package com.realimage.bean;

import com.opencsv.bean.CsvBindByPosition;

import lombok.Data;

@Data
public class Input {
	@CsvBindByPosition(position = 0)
	private String deliveryId;
	@CsvBindByPosition(position = 1)
	private int sizeOfDelivery;
	@CsvBindByPosition(position = 2)
	private String theatreId;

}
