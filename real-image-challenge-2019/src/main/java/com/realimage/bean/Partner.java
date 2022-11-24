package com.realimage.bean;

import com.opencsv.bean.CsvBindByName;

import lombok.Data;

@Data
public class Partner {
	@CsvBindByName(column = "Theatre")
	private String theatre;
	@CsvBindByName(column = "Size Slab (in GB)")
	private String sizeSlab;
	@CsvBindByName(column = "Minimum cost")
	private int minimumSlab;
	@CsvBindByName(column = "Cost Per GB")
	private int costPerGB;
	@CsvBindByName(column = "Partner ID")
	private String partnerId;

}
