package main

import (
	"log"
	"net/http"
	"shorturl"
)

func main() {

	handler := shorturl.RedirectHandler()

	log.Fatal(http.ListenAndServe(":8080", handler))
	//_ = shorturl.CreateURL("2a419415", "https://www.flipkart.com/mens-clothing/sports-wear/caps/pr?sid=2oq%2Cs9b%2C6gr%2Ch2i&p%5B%5D=facets.fulfilled_by%255B%255D%3DFlipkart%2BAssured&sort=popularity&fm=neo%2Fmerchandising&iid=M_e6ae5eb3-cfee-4714-b6e6-bf3eb2572d0a_14.A541W591A5GU&ssid=9r1kdnfgr40000001580014533197&otracker=hp_omu_Big%2BBrands%2Bon%2BBigger%2BDiscounts_3_7.dealCard.OMU_A541W591A5GU_5&otracker1=hp_omu_WHITELISTED_neo%2Fmerchandising_Big%2BBrands%2Bon%2BBigger%2BDiscounts_NA_dealCard_cc_3_NA_view-all_5&cid=A541W591A5GU")
}
