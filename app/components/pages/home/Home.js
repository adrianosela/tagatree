import React from 'react';
import { StyleSheet, View } from 'react-native';
import Map from '../../utils/Map';

export default class Home extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <View behavior='padding' style={styles.mapContainer}>
        <Map />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  mapContainer: {
    flex: 1,
    backgroundColor: '#81ecec'
  }
});
