package yelp

import "github.com/guregu/null"

//Float delegate for guregu/null
type Float struct {
	null.Float
}

//NewFloat Construction
func NewFloat(f float64, valid bool) Float { return Float{null.NewFloat(f, valid)} }

//FloatFrom Construction
func FloatFrom(f float64) Float { return Float{null.FloatFrom(f)} }

//FloatFromPtr Construction
func FloatFromPtr(f *float64) Float { return Float{null.FloatFromPtr(f)} }

//Int delegate for guregu/null
type Int struct {
	null.Int
}

//NewInt Construction
func NewInt(f int64, valid bool) Int { return Int{null.NewInt(f, valid)} }

//IntFrom Construction
func IntFrom(f int64) Int { return Int{null.IntFrom(f)} }

//IntFromPtr Construction
func IntFromPtr(f *int64) Int { return Int{null.IntFromPtr(f)} }

//Bool delegate for guregu/null
type Bool struct {
	null.Bool
}

//NewBool Construction
func NewBool(f bool, valid bool) Bool { return Bool{null.NewBool(f, valid)} }

//BoolFrom Construction
func BoolFrom(f bool) Bool { return Bool{null.BoolFrom(f)} }

//BoolFromPtr Construction
func BoolFromPtr(f *bool) Bool { return Bool{null.BoolFromPtr(f)} }
