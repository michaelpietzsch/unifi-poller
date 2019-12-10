package promunifi

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"golift.io/unifi"
)

type uap struct {
	// Ap Traffic Stats
	ApWifiTxDropped     *prometheus.Desc
	ApRxErrors          *prometheus.Desc
	ApRxDropped         *prometheus.Desc
	ApRxFrags           *prometheus.Desc
	ApRxCrypts          *prometheus.Desc
	ApTxPackets         *prometheus.Desc
	ApTxBytes           *prometheus.Desc
	ApTxErrors          *prometheus.Desc
	ApTxDropped         *prometheus.Desc
	ApTxRetries         *prometheus.Desc
	ApRxPackets         *prometheus.Desc
	ApRxBytes           *prometheus.Desc
	WifiTxAttempts      *prometheus.Desc
	MacFilterRejections *prometheus.Desc
	// VAP Stats
	VAPCcq                   *prometheus.Desc
	VAPMacFilterRejections   *prometheus.Desc
	VAPNumSatisfactionSta    *prometheus.Desc
	VAPAvgClientSignal       *prometheus.Desc
	VAPSatisfaction          *prometheus.Desc
	VAPSatisfactionNow       *prometheus.Desc
	VAPDNSAvgLatency         *prometheus.Desc
	VAPRxBytes               *prometheus.Desc
	VAPRxCrypts              *prometheus.Desc
	VAPRxDropped             *prometheus.Desc
	VAPRxErrors              *prometheus.Desc
	VAPRxFrags               *prometheus.Desc
	VAPRxNwids               *prometheus.Desc
	VAPRxPackets             *prometheus.Desc
	VAPTxBytes               *prometheus.Desc
	VAPTxDropped             *prometheus.Desc
	VAPTxErrors              *prometheus.Desc
	VAPTxPackets             *prometheus.Desc
	VAPTxPower               *prometheus.Desc
	VAPTxRetries             *prometheus.Desc
	VAPTxCombinedRetries     *prometheus.Desc
	VAPTxDataMpduBytes       *prometheus.Desc
	VAPTxRtsRetries          *prometheus.Desc
	VAPTxSuccess             *prometheus.Desc
	VAPTxTotal               *prometheus.Desc
	VAPTxGoodbytes           *prometheus.Desc
	VAPTxLatAvg              *prometheus.Desc
	VAPTxLatMax              *prometheus.Desc
	VAPTxLatMin              *prometheus.Desc
	VAPRxGoodbytes           *prometheus.Desc
	VAPRxLatAvg              *prometheus.Desc
	VAPRxLatMax              *prometheus.Desc
	VAPRxLatMin              *prometheus.Desc
	VAPWifiTxLatencyMovAvg   *prometheus.Desc
	VAPWifiTxLatencyMovMax   *prometheus.Desc
	VAPWifiTxLatencyMovMin   *prometheus.Desc
	VAPWifiTxLatencyMovTotal *prometheus.Desc
	VAPWifiTxLatencyMovCount *prometheus.Desc
	// Radio Stats
	RadioCurrentAntennaGain *prometheus.Desc
	RadioHt                 *prometheus.Desc
	RadioMaxTxpower         *prometheus.Desc
	RadioMinTxpower         *prometheus.Desc
	RadioNss                *prometheus.Desc
	RadioRadioCaps          *prometheus.Desc
	RadioTxPower            *prometheus.Desc
	RadioAstBeXmit          *prometheus.Desc
	RadioChannel            *prometheus.Desc
	RadioCuSelfRx           *prometheus.Desc
	RadioCuSelfTx           *prometheus.Desc
	RadioExtchannel         *prometheus.Desc
	RadioGain               *prometheus.Desc
	RadioNumSta             *prometheus.Desc
	RadioTxPackets          *prometheus.Desc
	RadioTxRetries          *prometheus.Desc
}

func descUAP(ns string) *uap {
	labelA := []string{"stat", "site_name", "name"} // stat + labels[1:]
	labelV := []string{"vap_name", "bssid", "radio", "radio_name", "essid", "usage", "site_name", "name"}
	labelR := []string{"radio_name", "radio", "site_name", "name"}
	return &uap{
		// 3x each - stat table: total, guest, user
		ApWifiTxDropped:     prometheus.NewDesc(ns+"stat_wifi_transmt_dropped_total", "Wifi Transmissions Dropped", labelA, nil),
		ApRxErrors:          prometheus.NewDesc(ns+"stat_receive_errors_total", "Receive Errors", labelA, nil),
		ApRxDropped:         prometheus.NewDesc(ns+"stat_receive_dropped_total", "Receive Dropped", labelA, nil),
		ApRxFrags:           prometheus.NewDesc(ns+"stat_receive_frags_total", "Received Frags", labelA, nil),
		ApRxCrypts:          prometheus.NewDesc(ns+"stat_receive_crypts_total", "Receive Crypts", labelA, nil),
		ApTxPackets:         prometheus.NewDesc(ns+"stat_transmit_packets_total", "Transmit Packets", labelA, nil),
		ApTxBytes:           prometheus.NewDesc(ns+"stat_transmit_bytes_total", "Transmit Bytes", labelA, nil),
		ApTxErrors:          prometheus.NewDesc(ns+"stat_transmit_errors_total", "Transmit Errors", labelA, nil),
		ApTxDropped:         prometheus.NewDesc(ns+"stat_transmit_dropped_total", "Transmit Dropped", labelA, nil),
		ApTxRetries:         prometheus.NewDesc(ns+"stat_retries_tx_total", "Transmit Retries", labelA, nil),
		ApRxPackets:         prometheus.NewDesc(ns+"stat_receive_packets_total", "Receive Packets", labelA, nil),
		ApRxBytes:           prometheus.NewDesc(ns+"stat_receive_bytes_total", "Receive Bytes", labelA, nil),
		WifiTxAttempts:      prometheus.NewDesc(ns+"stat_wifi_transmit_attempts_total", "Wifi Transmission Attempts", labelA, nil),
		MacFilterRejections: prometheus.NewDesc(ns+"stat_mac_filter_rejects_total", "MAC Filter Rejections", labelA, nil),
		// N each - 1 per Virtual AP (VAP)
		VAPCcq:                   prometheus.NewDesc(ns+"vap_ccq_ratio", "VAP Client Connection Quality", labelV, nil),
		VAPMacFilterRejections:   prometheus.NewDesc(ns+"vap_mac_filter_rejects_total", "VAP MAC Filter Rejections", labelV, nil),
		VAPNumSatisfactionSta:    prometheus.NewDesc(ns+"vap_satisfaction_stations", "VAP Number Satisifaction Stations", labelV, nil),
		VAPAvgClientSignal:       prometheus.NewDesc(ns+"vap_average_client_signal", "VAP Average Client Signal", labelV, nil),
		VAPSatisfaction:          prometheus.NewDesc(ns+"vap_satisfaction_ratio", "VAP Satisfaction", labelV, nil),
		VAPSatisfactionNow:       prometheus.NewDesc(ns+"vap_satisfaction_now_ratio", "VAP Satisfaction Now", labelV, nil),
		VAPDNSAvgLatency:         prometheus.NewDesc(ns+"vap_dns_latency_average_seconds", "VAP DNS Latency Average", labelV, nil),
		VAPRxBytes:               prometheus.NewDesc(ns+"vap_receive_bytes_total", "VAP Bytes Received", labelV, nil),
		VAPRxCrypts:              prometheus.NewDesc(ns+"vap_receive_crypts_total", "VAP Crypts Received", labelV, nil),
		VAPRxDropped:             prometheus.NewDesc(ns+"vap_receive_dropped_total", "VAP Dropped Received", labelV, nil),
		VAPRxErrors:              prometheus.NewDesc(ns+"vap_receive_errors_total", "VAP Errors Received", labelV, nil),
		VAPRxFrags:               prometheus.NewDesc(ns+"vap_receive_frags_total", "VAP Frags Received", labelV, nil),
		VAPRxNwids:               prometheus.NewDesc(ns+"vap_receive_nwids_total", "VAP Nwids Received", labelV, nil),
		VAPRxPackets:             prometheus.NewDesc(ns+"vap_receive_packets_total", "VAP Packets Received", labelV, nil),
		VAPTxBytes:               prometheus.NewDesc(ns+"vap_transmit_bytes_total", "VAP Bytes Transmitted", labelV, nil),
		VAPTxDropped:             prometheus.NewDesc(ns+"vap_transmit_dropped_total", "VAP Dropped Transmitted", labelV, nil),
		VAPTxErrors:              prometheus.NewDesc(ns+"vap_transmit_errors_total", "VAP Errors Transmitted", labelV, nil),
		VAPTxPackets:             prometheus.NewDesc(ns+"vap_transmit_packets_total", "VAP Packets Transmitted", labelV, nil),
		VAPTxPower:               prometheus.NewDesc(ns+"vap_transmit_power", "VAP Transmit Power", labelV, nil),
		VAPTxRetries:             prometheus.NewDesc(ns+"vap_transmit_retries_total", "VAP Retries Transmitted", labelV, nil),
		VAPTxCombinedRetries:     prometheus.NewDesc(ns+"vap_transmit_retries_combined_total", "VAP Retries Combined Transmitted", labelV, nil),
		VAPTxDataMpduBytes:       prometheus.NewDesc(ns+"vap_data_mpdu_transmit_bytes_total", "VAP Data MPDU Bytes Transmitted", labelV, nil),
		VAPTxRtsRetries:          prometheus.NewDesc(ns+"vap_transmit_rts_retries_total", "VAP RTS Retries Transmitted", labelV, nil),
		VAPTxSuccess:             prometheus.NewDesc(ns+"vap_transmit_success_total", "VAP Success Transmits", labelV, nil),
		VAPTxTotal:               prometheus.NewDesc(ns+"vap_transmit_total", "VAP Transmit Total", labelV, nil),
		VAPTxGoodbytes:           prometheus.NewDesc(ns+"vap_transmit_goodbyes", "VAP Goodbyes Transmitted", labelV, nil),
		VAPTxLatAvg:              prometheus.NewDesc(ns+"vap_transmit_latency_average_seconds", "VAP Latency Average Transmit", labelV, nil),
		VAPTxLatMax:              prometheus.NewDesc(ns+"vap_transmit_latency_maximum_seconds", "VAP Latency Maximum Transmit", labelV, nil),
		VAPTxLatMin:              prometheus.NewDesc(ns+"vap_transmit_latency_minimum_seconds", "VAP Latency Minimum Transmit", labelV, nil),
		VAPRxGoodbytes:           prometheus.NewDesc(ns+"vap_receive_goodbyes", "VAP Goodbyes Received", labelV, nil),
		VAPRxLatAvg:              prometheus.NewDesc(ns+"vap_receive_latency_average_seconds", "VAP Latency Average Receive", labelV, nil),
		VAPRxLatMax:              prometheus.NewDesc(ns+"vap_receive_latency_maximum_seconds", "VAP Latency Maximum Receive", labelV, nil),
		VAPRxLatMin:              prometheus.NewDesc(ns+"vap_receive_latency_minimum_seconds", "VAP Latency Minimum Receive", labelV, nil),
		VAPWifiTxLatencyMovAvg:   prometheus.NewDesc(ns+"vap_transmit_latency_moving_avg_seconds", "VAP Latency Moving Average Tramsit", labelV, nil),
		VAPWifiTxLatencyMovMax:   prometheus.NewDesc(ns+"vap_transmit_latency_moving_max_seconds", "VAP Latency Moving Maximum Tramsit", labelV, nil),
		VAPWifiTxLatencyMovMin:   prometheus.NewDesc(ns+"vap_transmit_latency_moving_min_seconds", "VAP Latency Moving Minimum Tramsit", labelV, nil),
		VAPWifiTxLatencyMovTotal: prometheus.NewDesc(ns+"vap_transmit_latency_moving_total", "VAP Latency Moving Total Tramsit", labelV, nil),
		VAPWifiTxLatencyMovCount: prometheus.NewDesc(ns+"vap_transmit_latency_moving_count", "VAP Latency Moving Count Tramsit", labelV, nil),
		// N each - 1 per Radio. 1-4 radios per AP usually
		RadioCurrentAntennaGain: prometheus.NewDesc(ns+"radio_current_antenna_gain", "Radio Current Antenna Gain", labelR, nil),
		RadioHt:                 prometheus.NewDesc(ns+"radio_ht", "Radio HT", labelR, nil),
		RadioMaxTxpower:         prometheus.NewDesc(ns+"radio_max_transmit_power", "Radio Maximum Transmit Power", labelR, nil),
		RadioMinTxpower:         prometheus.NewDesc(ns+"radio_min_transmit_power", "Radio Minimum Transmit Power", labelR, nil),
		RadioNss:                prometheus.NewDesc(ns+"radio_nss", "Radio Nss", labelR, nil),
		RadioRadioCaps:          prometheus.NewDesc(ns+"radio_caps", "Radio Capabilities", labelR, nil),
		RadioTxPower:            prometheus.NewDesc(ns+"radio_transmit_power", "Radio Transmit Power", labelR, nil),
		RadioAstBeXmit:          prometheus.NewDesc(ns+"radio_ast_be_xmit", "Radio AstBe Transmit", labelR, nil),
		RadioChannel:            prometheus.NewDesc(ns+"radio_channel", "Radio Channel", labelR, nil),
		RadioCuSelfRx:           prometheus.NewDesc(ns+"radio_channel_utilization_receive_ratio", "Radio Channel Utilization Receive", labelR, nil),
		RadioCuSelfTx:           prometheus.NewDesc(ns+"radio_channel_utilization_transmit_ratio", "Radio Channel Utilization Transmit", labelR, nil),
		RadioExtchannel:         prometheus.NewDesc(ns+"radio_ext_channel", "Radio Ext Channel", labelR, nil),
		RadioGain:               prometheus.NewDesc(ns+"radio_gain", "Radio Gain", labelR, nil),
		RadioNumSta:             prometheus.NewDesc(ns+"radio_stations", "Radio Total Station Count", append(labelR, "station_type"), nil),
		RadioTxPackets:          prometheus.NewDesc(ns+"radio_transmit_packets", "Radio Transmitted Packets", labelR, nil),
		RadioTxRetries:          prometheus.NewDesc(ns+"radio_transmit_retries", "Radio Transmit Retries", labelR, nil),
	}
}

func (u *promUnifi) exportUAP(r report, d *unifi.UAP) {
	labels := []string{d.Type, d.SiteName, d.Name}
	infoLabels := []string{d.Version, d.Model, d.Serial, d.Mac, d.IP, d.ID, d.Bytes.Txt, d.Uptime.Txt}
	u.exportUAPstats(r, labels, d.Stat.Ap, d.BytesD, d.TxBytesD, d.RxBytesD, d.BytesR)
	u.exportVAPtable(r, labels, d.VapTable)
	u.exportBYTstats(r, labels, d.TxBytes, d.RxBytes)
	u.exportSYSstats(r, labels, d.SysStats, d.SystemStats)
	u.exportSTAcount(r, labels, d.UserNumSta, d.GuestNumSta)
	u.exportRADtable(r, labels, d.RadioTable, d.RadioTableStats)
	r.send([]*metric{
		{u.Device.Info, gauge, 1.0, append(labels, infoLabels...)},
		{u.Device.Uptime, gauge, d.Uptime, labels},
	})
}

// udm doesn't have these stats exposed yet, so pass 2 or 6 metrics.
func (u *promUnifi) exportUAPstats(r report, labels []string, ap *unifi.Ap, bytes ...unifi.FlexInt) {
	if ap == nil {
		log.Println("ap was nil?!", labels[2])
		return
	}
	labelU := []string{"user", labels[1], labels[2]}
	labelG := []string{"guest", labels[1], labels[2]}
	r.send([]*metric{
		// ap only stuff.
		{u.Device.BytesD, counter, bytes[0], labels},   // not sure if these 3 Ds are counters or gauges.
		{u.Device.TxBytesD, counter, bytes[1], labels}, // not sure if these 3 Ds are counters or gauges.
		{u.Device.RxBytesD, counter, bytes[2], labels}, // not sure if these 3 Ds are counters or gauges.
		{u.Device.BytesR, gauge, bytes[3], labels},     // only UAP has this one, and those ^ weird.
		// user
		{u.UAP.ApWifiTxDropped, counter, ap.UserWifiTxDropped, labelU},
		{u.UAP.ApRxErrors, counter, ap.UserRxErrors, labelU},
		{u.UAP.ApRxDropped, counter, ap.UserRxDropped, labelU},
		{u.UAP.ApRxFrags, counter, ap.UserRxFrags, labelU},
		{u.UAP.ApRxCrypts, counter, ap.UserRxCrypts, labelU},
		{u.UAP.ApTxPackets, counter, ap.UserTxPackets, labelU},
		{u.UAP.ApTxBytes, counter, ap.UserTxBytes, labelU},
		{u.UAP.ApTxErrors, counter, ap.UserTxErrors, labelU},
		{u.UAP.ApTxDropped, counter, ap.UserTxDropped, labelU},
		{u.UAP.ApTxRetries, counter, ap.UserTxRetries, labelU},
		{u.UAP.ApRxPackets, counter, ap.UserRxPackets, labelU},
		{u.UAP.ApRxBytes, counter, ap.UserRxBytes, labelU},
		{u.UAP.WifiTxAttempts, counter, ap.UserWifiTxAttempts, labelU},
		{u.UAP.MacFilterRejections, counter, ap.UserMacFilterRejections, labelU},
		// guest
		{u.UAP.ApWifiTxDropped, counter, ap.GuestWifiTxDropped, labelG},
		{u.UAP.ApRxErrors, counter, ap.GuestRxErrors, labelG},
		{u.UAP.ApRxDropped, counter, ap.GuestRxDropped, labelG},
		{u.UAP.ApRxFrags, counter, ap.GuestRxFrags, labelG},
		{u.UAP.ApRxCrypts, counter, ap.GuestRxCrypts, labelG},
		{u.UAP.ApTxPackets, counter, ap.GuestTxPackets, labelG},
		{u.UAP.ApTxBytes, counter, ap.GuestTxBytes, labelG},
		{u.UAP.ApTxErrors, counter, ap.GuestTxErrors, labelG},
		{u.UAP.ApTxDropped, counter, ap.GuestTxDropped, labelG},
		{u.UAP.ApTxRetries, counter, ap.GuestTxRetries, labelG},
		{u.UAP.ApRxPackets, counter, ap.GuestRxPackets, labelG},
		{u.UAP.ApRxBytes, counter, ap.GuestRxBytes, labelG},
		{u.UAP.WifiTxAttempts, counter, ap.GuestWifiTxAttempts, labelG},
		{u.UAP.MacFilterRejections, counter, ap.GuestMacFilterRejections, labelG},
	})
}

// UAP VAP Table
func (u *promUnifi) exportVAPtable(r report, labels []string, vt unifi.VapTable) {
	// vap table stats
	for _, v := range vt {
		if !v.Up.Val {
			continue
		}
		labelV := []string{v.Name, v.Bssid, v.Radio, v.RadioName, v.Essid, v.Usage, labels[1], labels[2]}

		r.send([]*metric{
			{u.UAP.VAPCcq, gauge, float64(v.Ccq) / 1000.0, labelV},
			{u.UAP.VAPMacFilterRejections, counter, v.MacFilterRejections, labelV},
			{u.UAP.VAPNumSatisfactionSta, gauge, v.NumSatisfactionSta, labelV},
			{u.UAP.VAPAvgClientSignal, gauge, v.AvgClientSignal.Val, labelV},
			{u.UAP.VAPSatisfaction, gauge, v.Satisfaction.Val / 100.0, labelV},
			{u.UAP.VAPSatisfactionNow, gauge, v.SatisfactionNow.Val / 100.0, labelV},
			{u.UAP.VAPDNSAvgLatency, gauge, v.DNSAvgLatency.Val / 1000, labelV},
			{u.UAP.VAPRxBytes, counter, v.RxBytes, labelV},
			{u.UAP.VAPRxCrypts, counter, v.RxCrypts, labelV},
			{u.UAP.VAPRxDropped, counter, v.RxDropped, labelV},
			{u.UAP.VAPRxErrors, counter, v.RxErrors, labelV},
			{u.UAP.VAPRxFrags, counter, v.RxFrags, labelV},
			{u.UAP.VAPRxNwids, counter, v.RxNwids, labelV},
			{u.UAP.VAPRxPackets, counter, v.RxPackets, labelV},
			{u.UAP.VAPTxBytes, counter, v.TxBytes, labelV},
			{u.UAP.VAPTxDropped, counter, v.TxDropped, labelV},
			{u.UAP.VAPTxErrors, counter, v.TxErrors, labelV},
			{u.UAP.VAPTxPackets, counter, v.TxPackets, labelV},
			{u.UAP.VAPTxPower, gauge, v.TxPower, labelV},
			{u.UAP.VAPTxRetries, counter, v.TxRetries, labelV},
			{u.UAP.VAPTxCombinedRetries, counter, v.TxCombinedRetries, labelV},
			{u.UAP.VAPTxDataMpduBytes, counter, v.TxDataMpduBytes, labelV},
			{u.UAP.VAPTxRtsRetries, counter, v.TxRtsRetries, labelV},
			{u.UAP.VAPTxTotal, counter, v.TxTotal, labelV},
			{u.UAP.VAPTxGoodbytes, counter, v.TxTCPStats.Goodbytes, labelV},
			{u.UAP.VAPTxLatAvg, gauge, v.TxTCPStats.LatAvg.Val / 1000, labelV},
			{u.UAP.VAPTxLatMax, gauge, v.TxTCPStats.LatMax.Val / 1000, labelV},
			{u.UAP.VAPTxLatMin, gauge, v.TxTCPStats.LatMin.Val / 1000, labelV},
			{u.UAP.VAPRxGoodbytes, counter, v.RxTCPStats.Goodbytes, labelV},
			{u.UAP.VAPRxLatAvg, gauge, v.RxTCPStats.LatAvg.Val / 1000, labelV},
			{u.UAP.VAPRxLatMax, gauge, v.RxTCPStats.LatMax.Val / 1000, labelV},
			{u.UAP.VAPRxLatMin, gauge, v.RxTCPStats.LatMin.Val / 1000, labelV},
			{u.UAP.VAPWifiTxLatencyMovAvg, gauge, v.WifiTxLatencyMov.Avg.Val / 1000, labelV},
			{u.UAP.VAPWifiTxLatencyMovMax, gauge, v.WifiTxLatencyMov.Max.Val / 1000, labelV},
			{u.UAP.VAPWifiTxLatencyMovMin, gauge, v.WifiTxLatencyMov.Min.Val / 1000, labelV},
			{u.UAP.VAPWifiTxLatencyMovTotal, counter, v.WifiTxLatencyMov.Total, labelV},      // not sure if gauge or counter.
			{u.UAP.VAPWifiTxLatencyMovCount, counter, v.WifiTxLatencyMov.TotalCount, labelV}, // not sure if gauge or counter.
		})
	}
}

// UAP Radio Table
func (u *promUnifi) exportRADtable(r report, labels []string, rt unifi.RadioTable, rts unifi.RadioTableStats) {
	// radio table
	for _, p := range rt {
		labelR := []string{p.Name, p.Radio, labels[1], labels[2]}
		labelRUser := append(labelR, "user")
		labelRGuest := append(labelR, "guest")
		r.send([]*metric{
			{u.UAP.RadioCurrentAntennaGain, gauge, p.CurrentAntennaGain, labelR},
			{u.UAP.RadioHt, gauge, p.Ht, labelR},
			{u.UAP.RadioMaxTxpower, gauge, p.MaxTxpower, labelR},
			{u.UAP.RadioMinTxpower, gauge, p.MinTxpower, labelR},
			{u.UAP.RadioNss, gauge, p.Nss, labelR},
			{u.UAP.RadioRadioCaps, gauge, p.RadioCaps, labelR},
		})

		// combine radio table with radio stats table.
		for _, t := range rts {
			if t.Name != p.Name {
				continue
			}
			r.send([]*metric{
				{u.UAP.RadioTxPower, gauge, t.TxPower, labelR},
				{u.UAP.RadioAstBeXmit, gauge, t.AstBeXmit, labelR},
				{u.UAP.RadioChannel, gauge, t.Channel, labelR},
				{u.UAP.RadioCuSelfRx, gauge, t.CuSelfRx.Val / 100.0, labelR},
				{u.UAP.RadioCuSelfTx, gauge, t.CuSelfTx.Val / 100.0, labelR},
				{u.UAP.RadioExtchannel, gauge, t.Extchannel, labelR},
				{u.UAP.RadioGain, gauge, t.Gain, labelR},
				{u.UAP.RadioNumSta, gauge, t.GuestNumSta, labelRGuest},
				{u.UAP.RadioNumSta, gauge, t.UserNumSta, labelRUser},
				{u.UAP.RadioTxPackets, gauge, t.TxPackets, labelR},
				{u.UAP.RadioTxRetries, gauge, t.TxRetries, labelR},
			})
			break
		}
	}
}
