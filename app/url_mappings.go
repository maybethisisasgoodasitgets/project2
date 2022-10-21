package app

import "porject2/controllers/ping"




func mapUrls() {

    router.GET("/ping", ping.Ping)

}