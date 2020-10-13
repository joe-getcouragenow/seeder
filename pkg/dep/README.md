# deps

These are all the thngs that the build needs on the OS

Protobuf, Go, and GPRC

Windows hover uses candle and light.
download and unpack from here.
https://github.com/wixtoolset/wix3/releases/tag/wix3112rtm

Darwin and Linux uses shittty nodejs like.
I made a quasi PR: https://github.com/go-flutter-desktop/go-flutter/issues/521

Flutter needs Android SDK and NDK
- For CI it gets installed anyway
- But for Dev it does not.
- here is a bash installer: https://github.com/google/mediapipe/blob/master/setup_android_sdk_and_ndk.sh
    - we can adapt for Mac
