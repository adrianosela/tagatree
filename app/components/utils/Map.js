import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import MapView from 'react-native-maps';

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
            <MapView
              style={styles.map}
              region={{
                latitude: location.coordinates[1],
                longitude: location.coordinates[0],
                latitudeDelta: 0.01,
                longitudeDelta: 0.01
              }}
	    >
              <MapView.Marker coordinate={{latitude: location.coordinates[1], longitude: location.coordinates[0]}} />
	    </MapView>
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
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  titleText: {
    fontSize: 30,
    marginTop: 50
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
  },
  map: {
    ...StyleSheet.absoluteFill
  }
});
