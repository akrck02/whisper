package service

import (
	"fmt"
	"log"

	"github.com/pion/webrtc/v4"
)

func StartRTC() {

	// Everything below is the Pion WebRTC API!
	// We define the configuration for ICE
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Peer connection created successfully!")

	// Close the peer connection when you're done with it
	defer func() {
		if err := peerConnection.Close(); err != nil {
			fmt.Printf("Error closing peer connection: %s", err)
		}
	}()

	// Code for handling ICE candidates and sending/receiving media streams will go here

	// Create a new track
	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeVP8,
	}, "video", "pion")
	if err != nil {
		panic(err)
	}

	rtpSender, err := peerConnection.AddTrack(track)
	if err != nil {
		panic(err)
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by the interceptors.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			i, _, rtcpErr := rtpSender.Read(rtcpBuf)
			if rtcpErr != nil {
				return
			}

			fmt.Println("RTCP packet received!")
			_ = i
		}
	}()

	// Handle incoming media streams
	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		// Print incoming stream's info
		fmt.Printf("Got track: %+v", track)

		// Read incoming packets
		b := make([]byte, 1500)
		for {
			i, _, err := track.Read(b)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Got packet with length: %d", i)
			// Process the packet here
		}
	})

	// Wait for user input to keep the connection alive
	fmt.Println("Press Ctrl+C to exit")
	select {}
}
