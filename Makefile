#

txt=delegated-apnic.txt
split=awksplit
gen=awkgen

json=ipnet-cn.json
gosrc=ipnet.go

$(gosrc): $(txt)
	cat ${txt} | awk -f ${split} | sort -k1,1 -n | (apnic_target=go awk -f ${gen}) > $@

$(json): $(txt)
	cat ${txt} | awk -f ${split} | sort -k1,1 -n | awk -f ${gen} > $@

$(txt):
	curl "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest" > $@

clean:
	rm -f ${json} ${gosrc}

nuke: clean
	rm -f ${txt}
