package tool

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func download() {
	dir := "./AAA"
	os.MkdirAll(dir, os.ModeDir)
	urls := []string{
		"https://www.scnet.cn/ui/chatbot/_nuxt/D2UNSEaE.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CfIEWrDu.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/2ruCGqZH.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BmfIzHrd.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BxkM9zWh.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/C1bfHdrV.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/wGPHhbaM.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BFUxVpKz.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/QylXqPc7.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/qzc_tsZ9.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BDrZ2-3S.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/k8a50bkU.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DJrMQMoq.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BaUGqCO4.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/5kbUw4KD.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/1h0P1dVm.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D2qaP0EU.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DBlgUQyS.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DUC-38up.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CnnhLmvs.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Cmpz8WtF.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/lVrDEbk_.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CEi5DQ6j.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/WftPVMHc.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/wFEio8b6.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Yn2BQbom.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CTMT6Tlx.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D3L284zb.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Cix1fYFN.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/ZOZghsSW.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BXwtZ7iO.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CM_jh9vA.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/5bihd-W0.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/nWa69ocm.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DfEOvBz1.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DUoMKKtV.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/hs2IdhKk.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/xFfyNXWI.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BZPocrqR.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BSWU5ELP.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CeAX10gq.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/lijyi5DM.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D5HqFXIy.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Wq0dG51x.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/FBkZzy2B.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Ca_7UZP3.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/IAAswi8U.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CY3RDbwi.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CBXIumtt.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/2lziiFmm.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CRRCJcXR.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/55xxL4Ln.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/iJpoKaFc.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/yJDoF_Pv.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DVfDLyVg.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BSPShLeb.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/juUdrNN7.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/59H9CEe_.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/dya6X0Yb.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DSRKrrXi.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/_H57VE8H.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/xqf1j9LT.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/94KM8GKN.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/oBs6DRoB.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DB1Ud2Wy.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/B2HqeDee.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Bs9wJ_za.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/zRyxzPTX.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BHU3785f.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/C35FgnSH.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Bryeh17Y.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D3xHyQcA.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/2ADw5Ftp.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DJBVk8Hd.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Dtjc5phW.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CC3E01oj.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/Dhd2JWle.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CTtZtRSL.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CkHfQAAA.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BPhmSymZ.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D0tQaE1T.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CVqUeTsp.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/D0sCm7f7.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DezyIfNm.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DMAR3mfp.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CDf7lZEH.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DNrjPWqC.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DlJwyu6q.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CknpES6Q.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/DDHQFboy.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/CMoyv-5G.js",
		"https://www.scnet.cn/ui/chatbot/_nuxt/BfnZmUPj.js",
	}
	for _, u := range urls {
		fmt.Println(u)
		resp, err := http.Get(u)
		if err != nil {
			fmt.Println(err)
			return
		}
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		dd := strings.Split(u, "/")
		fileName := dd[len(dd)-1]

		fs, err := os.Create(path.Join(dir, fileName))
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = fs.Write(bs)
		if err != nil {
			fmt.Println(err)
			return
		}
		fs.Close()
	}
}

func Filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, v := range data {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
