import React from 'react';
import { StyleSheet, Text, View } from 'react-native';

export default class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      location: null,
    };
  }

  componentDidMount() {
    navigator.geolocation.getCurrentPosition(
      (position) => {
	this.setState({
	  loading: false,
	  location: {
	    coordinates: [position.coords.longitude, position.coords.latitude]
	  }
	});
      },
      (error) => {
	console.log(error);
      }
    );
  }

  render() {
    const { loading, location } = this.state;

    return (
      <View style={styles.appContainer}>
        { loading ? (
          <View style={styles.loadingContainer}>
            <Text style={styles.titleText}>Locating You</Text>
          </View>
          ) : (
          <View style={styles.bodyContainer}>
            <Text style={styles.titleText}>Your Location:</Text>
            <Text style={styles.coordsText}>Latitude: {location.coordinates[1]}</Text>
            <Text style={styles.coordsText}>Longitude: {location.coordinates[0]}</Text>
          </View>
        )}
      </View>
    );
  }

}

const styles = StyleSheet.create({
  appContainer: {
    flex: 1,
    backgroundColor: '#FFF'
  },
  loadingContainer: {
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#00FF00'
  },
  titleText: {
    fontSize: 40,
    marginBottom: 50
  },
  coordsText: {
    marginBottom: 20,
    fontSize: 30,
    color: '#000000'
  },
  bodyContainer: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#90FF28'
  }
});
