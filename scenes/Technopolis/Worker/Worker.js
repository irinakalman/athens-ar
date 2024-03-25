import React, { useState, useEffect } from 'react';
import {
  ViroARScene,
  Viro3DObject,
  ViroARPlaneSelector,
  ViroTrackingStateConstants,
  ViroAmbientLight
} from '@viro-community/react-viro';
import { StyleSheet } from "react-native";

var styles = StyleSheet.create({
  f1: { flex: 1 },
  helloWorldTextStyle: {
    fontFamily: "Arial",
    fontSize: 30,
    color: "#ffffff",
    textAlignVertical: "center",
    textAlign: "center",
  },
});

const Worker = () => {
  const [text, setText] = useState("Initializing AR...");

  function onInitialized(state: any, reason: ViroTrackingReason) {
    console.log("onInitialized", state, reason);
    if (state === ViroTrackingStateConstants.TRACKING_NORMAL) {
      setText("Hello World!");
    } else if (state === ViroTrackingStateConstants.TRACKING_UNAVAILABLE) {
      // Handle loss of tracking
    }
  }

  return (
    <ViroARScene onTrackingUpdated={onInitialized}>
      <ViroAmbientLight color="#ffffff" intensity={200}/>
      <Viro3DObject
        source={require('./Duck.glb')}
        position={[0, 0, -1]}
        scale={[1, 1, 1]}
        type="GLB"
      />
    </ViroARScene>
  );
};

export default Worker;
