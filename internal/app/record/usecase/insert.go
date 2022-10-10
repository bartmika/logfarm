package usecase

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2/format"

	r_d "github.com/bartmika/logfarm/internal/domain/record"
)

func (uc recordUsecase) InsertLogParts(ctx context.Context, logParts format.LogParts) (record *r_d.Record, err error) {
	client, clientOK := logParts["client"].(string)
	if !clientOK {
		client = ""
	}
	content, contentOK := logParts["content"].(string)
	if !contentOK {
		content = ""
	}
	facility, facilityOK := logParts["facility"].(int)
	if !facilityOK {
		facility = 0
	}
	hostname, hostnameOK := logParts["hostname"].(string)
	if !hostnameOK {
		hostname = ""
	}
	priority, priorityOK := logParts["priority"].(int)
	if !priorityOK {
		priority = 0
	}
	severity, severityOK := logParts["severity"].(int)
	if !severityOK {
		severity = 0
	}
	tag, tagOK := logParts["tag"].(string)
	if !tagOK {
		tag = ""
	}
	timestamp, timestampOK := logParts["timestamp"].(time.Time)
	if !timestampOK {
		timestamp = time.Now()
	}
	tls_peer, tls_peerOK := logParts["tls_peer"].(string)
	if !tls_peerOK {
		tls_peer = ""
	}

	r := &r_d.Record{
		ID:        uc.UUID.NewUUID(),
		Client:    client,
		Content:   content,
		Facility:  facility,
		Hostname:  hostname,
		Priority:  priority,
		Severity:  severity,
		Tag:       tag,
		Timestamp: timestamp,
		TLSPeer:   tls_peer,
	}

	err = uc.RecordRepo.Insert(ctx, r)
	if err != nil {
		return nil, err
	}

	log.Info().
		Str("Function", "InsertLogParts").
		Str("Client", client).
		Str("Content", content).
		Int("Facility", facility).
		Str("Hostname", hostname).
		Int("Priority", priority).
		Int("Severity", severity).
		Str("Tag", tag).
		Time("Timestamp", timestamp).
		Str("TLSPeer", tls_peer).
		Msg("Record created")

	return r, nil
}
