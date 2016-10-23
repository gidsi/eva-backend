package main

type Feed struct {
        Type string `json:"type"`
        Url string `json:"url"`
}

type Location struct {
        Address string `json:"address"`
        Lon float32 `json:"lon"`
        Lat float32 `json:"lat"`
}

type Contact struct {
        Twitter string `json:"twitter"`
        Phone string `json:"phone"`
        Irc string `json:"irc"`
        Email string `json:"email"`
        Ml string `json:"ml"`
        IssueMail string `json:"issue_mail"`
}

type State struct {
        Open bool `json:"open"`
}

type Feeds struct {
        Blog Feed `json:"blog"`
        Wiki Feed `json:"wiki"`
        Calendar Feed `json:"calendar"`
}

type SpaceData struct {
        Api string `json:"api"`
        Space string `json:"space"`
        Logo string `json:"logo"`
        Url string `json:"url"`
        Location Location `json:"location"`
        Contact Contact `json:"contact"`
        IssueReportChannels []string `json:"issue_report_channels"`
        State State `json:"state"`
        Feeds Feeds `json:"feeds"`
        Ext_ccc string `json:"ext_ccc"`
}