package main

import (
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := NewForm(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	form = NewForm(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form post does not have required fields, when it does")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Valid returns false and it should be true when calling Check()")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")

	if len(s) == 0 {
		t.Error("should have an error returned from Get, but does not")
	}

	s = form.Errors.Get("whatever")
	if len(s) != 0 {
		t.Error("should not have an error, but got one")
	}
}
