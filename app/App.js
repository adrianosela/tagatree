import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Tree from './components/Tree'

export default class App extends React.Component {

  state = {
    isLoading: true,
    location: {
      coordinates: [null, null]
    }
  };

  componentDidMount() {
    navigator.geolocation.getCurrentPosition(
      position => {
	this.setState({
	  isLoading: false,
	  location: {
	    coordinates: [position.coords.longitude, position.coords.latitude]
	  }
	});
      },
      error => {
	console.log(error); // FIXME
      }
    );
  }
  
  render() {
    const { isLoading, location } = this.state

    return (
      <View style={styles.appContainer}>
        {isLoading ? (
          <View style={styles.loadingContainer}>
            <Text stlye={styles.loadingText}>Locating You</Text> 
          </View>
          ) : (
          <View style={styles.bodyContainer}>
            <Tree species='Sprucea' location={location}/>
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
    backgroundColor: '#00FF00'
  },
  loadingText: {
    fontSize: 200
  },
  bodyContainer: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#00FF8F'
  }
});
