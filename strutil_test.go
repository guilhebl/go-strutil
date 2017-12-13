package strutil

import (
	"testing"
)

func TestFilterHtmlTags(t *testing.T) {
	s := `
<p>Links:</p><ul><li><a href="foo">Foo</a><li>
<a href="/bar/baz">BarBaz</a></ul><span>TEXT <b>I</b> WANT</span>
<script type='text/javascript'>
/* <![CDATA[ */
var post_notif_widget_ajax_obj = {"ajax_url":"http:\/\/site.com\/wp-admin\/admin-ajax.php","nonce":"9b8270e2ef","processing_msg":"Processing..."};
/* ]]> */
</script>`

	r := FilterHtmlTags(s)

	if r != "Links:FooBarBazTEXTIWANT" {
		t.Error("Error while filtering HTML string")
	}

	s2 := `
	<!DOCTYPE html PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<HTML>
	   <HEAD>
		  <TITLE>
			 A Small Hello
		  </TITLE>
	   </HEAD>
	<BODY>
	<H1>Hi</H1><P>This is very minimal "hello world" HTML document.</P><div><h3>content</h3></div>
	</BODY>
	</HTML>`
	r2 := FilterHtmlTags(s2)

	if r2 != "A Small HelloHiThis is very minimal \"hello world\" HTML document.content" {
		t.Error("Error while filtering HTML string")
	}
}

func TestTrimSuffix(t *testing.T) {

	s := "tree+abc+def+gef+"

	s = TrimSuffix(s, "+")

	if s != "tree+abc+def+gef" {
		t.Errorf("Incorrect trimmed suffix for %s", s)
	}
}