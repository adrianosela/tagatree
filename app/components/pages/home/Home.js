import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Tree from '../../objects/Tree'

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
            <Text stlye={styles.loadingText}>Locating You</Text>
          </View>
          ) : (
          <View style={styles.bodyContainer}>
            <Tree species='Spruce' location={location}/>
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
