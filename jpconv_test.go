package jpconv

import (
	"testing"
)

func TestHiraganaToKatakana(t *testing.T) {
	var result string
	result = HiraganaToKatakana("ひらがなをカタカナにへんかんします")
	if result != "ヒラガナヲカタカナニヘンカンシマス" {
		t.Errorf("fail! %s", result)
	}
	result = HiraganaToKatakana("ゔぁいおりん")
	if result != "ヴァイオリン" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkHiraganaToKatakana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HiraganaToKatakana("ひらがなをカタカナにへんかんします")
	}
}

func TestKatakanaToHiragana(t *testing.T) {
	var result string
	result = KatakanaToHiragana("カタカナヲひらがなニヘンカンシマス")
	if result != "かたかなをひらがなにへんかんします" {
		t.Errorf("fail! %s", result)
	}
	result = KatakanaToHiragana("ヴァイオリン")
	if result != "ゔぁいおりん" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkKatakanaToHiragana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KatakanaToHiragana("カタカナヲひらがなニヘンカンシマス")
	}
}

func TestHankakuToZenkaku(t *testing.T) {
	var result string
	result = HankakuToZenkaku("ﾊﾝｶｸｦｾﾞﾝｶｸﾆﾍﾝｶﾝｼﾏｽ")
	if result != "ハンカクヲゼンカクニヘンカンシマス" {
		t.Error("fail!")
	}
}

func BenchmarkHankakuToZenkaku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HankakuToZenkaku("ﾊﾝｶｸｦｾﾞﾝｶｸﾆﾍﾝｶﾝｼﾏｽ")
	}
}

func TestZenkakuToHankaku(t *testing.T) {
	var result string
	result = ZenkakuToHankaku("ゼンカクヲハンカクニヘンカンシマス")
	if result != "ｾﾞﾝｶｸｦﾊﾝｶｸﾆﾍﾝｶﾝｼﾏｽ" {
		t.Error("fail!")
	}
}

func BenchmarkZenkakuToHankaku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZenkakuToHankaku("ゼンカクヲハンカクニヘンカンシマス")
	}
}
