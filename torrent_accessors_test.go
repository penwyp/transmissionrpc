package transmissionrpc

import (
	"encoding/json"
	"log"
	"testing"
)

func Test_xx(t *testing.T) {

	body := `
{
                "activityDate": 1668231638,
                "addedDate": 1665413621,
                "bandwidthPriority": 0,
                "comment": "",
                "corruptEver": 0,
                "creator": "uTorrent/2210",
                "dateCreated": 1452811090,
                "desiredAvailable": 0,
                "doneDate": 0,
                "downloadDir": "/volume2/pt/\u5176\u4ed6/\u5b66\u4e60",
                "downloadLimit": 1,
                "downloadLimited": false,
                "downloadedEver": 0,
                "editDate": 0,
                "error": 2,
                "errorString": "Tracker HTTP response 502 (Bad Gateway",
                "eta": -1,
                "etaIdle": -1,
                "fileStats": [
                    {
                        "bytesCompleted": 42087610,
                        "priority": 0,
                        "wanted": true
                    }
                ],
                "files": [
                    {
                        "bytesCompleted": 42087610,
                        "length": 42087610,
                        "name": "\u300a\u4e09\u8054\u751f\u6d3b\u5468\u520a\u300b2016\u5e74\u7b2c01\u671f\uff08\u672a\u6765\u7b80\u53f2\uff09.PDF"
                    }
                ],
                "hashString": "00155acc9a607671ecb63fe94e20a5202881c35f",
                "haveUnchecked": 0,
                "haveValid": 42087610,
                "honorsSessionLimits": true,
                "id": 4528,
                "isFinished": false,
                "isPrivate": true,
                "isStalled": false,
                "labels": [],
                "leftUntilDone": 0,
                "magnetLink": "magnet:?xt=urn:btih:00155acc9a607671ecb63fe94e20a5202881c35f&dn=%E3%80%8A%E4%B8%89%E8%81%94%E7%94%9F%E6%B4%BB%E5%91%A8%E5%88%8A%E3%80%8B2016%E5%B9%B4%E7%AC%AC01%E6%9C%9F%EF%BC%88%E6%9C%AA%E6%9D%A5%E7%AE%80%E5%8F%B2%EF%BC%89.PDF&tr=https%3A%2F%2Ftracker.m-team.cc%2Fannounce.php%3Fpasskey%3D1b376abe27739b012fb9b5a226452d13",
                "manualAnnounceTime": -1,
                "maxConnectedPeers": 50,
                "metadataPercentComplete": 1,
                "name": "\u300a\u4e09\u8054\u751f\u6d3b\u5468\u520a\u300b2016\u5e74\u7b2c01\u671f\uff08\u672a\u6765\u7b80\u53f2\uff09.PDF",
                "peer-limit": 50,
                "peers": [],
                "peersConnected": 0,
                "peersFrom": {
                    "fromCache": 0,
                    "fromDht": 0,
                    "fromIncoming": 0,
                    "fromLpd": 0,
                    "fromLtep": 0,
                    "fromPex": 0,
                    "fromTracker": 0
                },
                "peersGettingFromUs": 0,
                "peersSendingToUs": 0,
                "percentDone": 1,
                "pieceCount": 643,
                "pieceSize": 65536,
                "pieces": "///////////////////////////////////////////////////////////////////////////////////////////////////////////g",
                "priorities": [
                    0
                ],
                "queuePosition": 4527,
                "rateDownload": 0,
                "rateUpload": 0,
                "recheckProgress": 0,
                "secondsDownloading": 1154,
                "secondsSeeding": 1267393,
                "seedIdleLimit": 30,
                "seedIdleMode": 0,
                "seedRatioLimit": 2,
                "seedRatioMode": 0,
                "sizeWhenDone": 42087610,
                "startDate": 1673714118,
                "status": 6,
                "torrentFile": "/config/torrents/00155acc9a607671ecb63fe94e20a5202881c35f.torrent",
                "totalSize": 42087610,
                "trackerStats": [
                    {
                        "announce": "https://tracker.m-team.cc/announce.php?passkey=1b376abe27739b012fb9b5a226452d13",
                        "announceState": 1,
                        "downloadCount": -1,
                        "hasAnnounced": true,
                        "hasScraped": true,
                        "host": "tracker.m-team.cc:443",
                        "id": 4536,
                        "isBackup": false,
                        "lastAnnouncePeerCount": 0,
                        "lastAnnounceResult": "Tracker HTTP response 502 (Bad Gateway",
                        "lastAnnounceStartTime": 0,
                        "lastAnnounceSucceeded": false,
                        "lastAnnounceTime": 1673714327,
                        "lastAnnounceTimedOut": false,
                        "lastScrapeResult": "Tracker HTTP response 502 (Bad Gateway)",
                        "lastScrapeStartTime": 0,
                        "lastScrapeSucceeded": false,
                        "lastScrapeTime": 1673714511,
                        "lastScrapeTimedOut": false,
                        "leecherCount": -1,
                        "nextAnnounceTime": 1673715286,
                        "nextScrapeTime": 1673716330,
                        "scrape": "https://tracker.m-team.cc/scrape.php?passkey=1b376abe27739b012fb9b5a226452d13",
                        "scrapeState": 1,
                        "seederCount": -1,
                        "sitename": "m-team",
                        "tier": 0
                    }
                ],
                "trackers": [
                    {
                        "announce": "https://tracker.m-team.cc/announce.php?passkey=1b376abe27739b012fb9b5a226452d13",
                        "id": 4536,
                        "scrape": "https://tracker.m-team.cc/scrape.php?passkey=1b376abe27739b012fb9b5a226452d13",
                        "sitename": "m-team",
                        "tier": 0
                    }
                ],
                "uploadLimit": 100,
                "uploadLimited": false,
                "uploadRatio": 1.0521,
                "uploadedEver": 44282145,
                "wanted": [
                    true
                ],
                "webseeds": [],
                "webseedsSendingToUs": 0
            }
`

	item := Torrent{}

	uErr := json.Unmarshal([]byte(body), &item)
	if uErr != nil {
		log.Fatal(uErr)
	}
	log.Println(item)
}
