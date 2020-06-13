import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const Tree = ({ species, location }) => {
  return (
    <View style={styles.treeContainer}>
      <View style={styles.bodyContainer}>
        <Text style={styles.speciesText}>{species}</Text>
        <Text style={styles.locationText}>Latitude: {location.coordinates[1]}</Text>
        <Text style={styles.locationText}>Longitude: {location.coordinates[0]}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  treeContainer: {
    flex: 1
  },
  bodyContainer: {
    flex: 2,
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
    paddingLeft: 25,
    marginTop: 100
  },
  speciesText: {
    fontSize: 65,
    color: '#fff'
  },
  locationText: {
    fontSize: 60,
    color: '#000000'
  }
});

export default Tree;
