package main

import (
	"fmt"
	"os"
	"shorturl"
)

func main() {

	//handler := shorturl.RedirectHandler()

	//log.Fatal(http.ListenAndServe(":8080", handler))
	urlArr := [10]string{"https://github.com/farzamalam", "https://www.youtube.com/", "https://www.google.com/", "https://www.amazon.com/BioShock-Daddy-Backpack-not-machine-specific/dp/B01MSUOG4A?pf_rd_p=04fcd686-758d-4aa8-85e7-3b200e207bde&pd_rd_wg=vd8NS&pf_rd_r=G7NE5EDAZ41MGNC582AR&ref_=pd_gw_unk&pd_rd_w=UeVzd&pd_rd_r=f194abae-09a8-44c2-9ca4-b1f9b4470623", "https://www.kaggle.com/dejavu23/titanic-survival-seaborn-and-ensembles", "https://github.com/gophercises/urlshort/blob/master/students/kalexmills/main/main.go", "flipkart.com/mens-clothing/sports-wear/caps/pr?sid=2oq%2Cs9b%2C6gr%2Ch2i&p%5B%5D=facets.fulfilled_by%255B%255D%3DFlipkart%2BAssured&sort=popularity&fm=neo%2Fmerchandising&iid=M_e6ae5eb3-cfee-4714-b6e6-bf3eb2572d0a_14.A541W591A5GU&ssid=9r1kdnfgr40000001580014533197&otracker=hp_omu_Big%2BBrands%2Bon%2BBigger%2BDiscounts_3_7.dealCard.OMU_A541W591A5GU_5&otracker1=hp_omu_WHITELISTED_neo%2Fmerchandising_Big%2BBrands%2Bon%2BBigger%2BDiscounts_NA_dealCard_cc_3_NA_view-all_5&cid=A541W591A5GU", "https://www.myntra.com/blazers/jack--jones/jack--jones-black-premium-slim-fit-solid-single-breasted-casual-corduroy-blazer/8860551/buy", "https://www.udemy.com/", "https://tour.golang.org/list"}
	for i := 0; i < len(urlArr); i++ {
		srcHash := shorturl.GenerateHash(urlArr[i])[:8]
		err := shorturl.CreateURL(srcHash, urlArr[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while storing key value: %v", err)
		}
		fmt.Println(srcHash, urlArr[i])
	}
}
