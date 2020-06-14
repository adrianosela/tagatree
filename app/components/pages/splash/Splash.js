import React from 'react';
import { StyleSheet, View, StatusBar, Animated, ActivityIndicator } from 'react-native';
import jwtDecode from 'jwt-decode';
import PropTypes from 'prop-types';
import AsyncStorageManager from '../../storage/AsyncStorageManager';
import Logo from '../../utils/Logo';

export default class Splash extends React.Component {
  propTypes = {
    navigation: PropTypes.shape({
      navigate: PropTypes.func.isRequired,
    }).isRequired,
  }

  constructor(props) {
    super(props);

    this.state = {
      logoWidth: new Animated.Value(80),
      logoOpacity: new Animated.Value(0),
      sloganOpacity: new Animated.Value(0),
      showLoading: false,
    };
  }

  componentDidMount() {
    this.animateLogo().then(() => {
      return AsyncStorageManager.getInstance().getUserToken();
    }).then(userToken => {
      if (!userToken) {
        throw new Error('Empty token');
      }

      const token = userToken.split(' ')[1];
      const decoded = jwtDecode(token);
      const exp = new Date(decoded.exp * 1000);

      if (exp > new Date()) {
        // Valid token, skip login
        this.props.navigation.navigate('Home');
      } else {
        // Not valid
        this.props.navigation.navigate('Login');
      }
    }).catch(error => {
      console.log(error);
      this.props.navigation.navigate('Login');
    });
  }

  animateLogo() {
    return new Promise((resolve) => {
      Animated.sequence([
        Animated.timing(
          this.state.logoWidth,
          {
            toValue: 260,
            duration: 700,
          }
        ),
        Animated.timing(
          this.state.logoOpacity,
          {
            toValue: 1.0,
            duration: 700,
          }
        ),
        Animated.timing(
          this.state.sloganOpacity,
          {
            toValue: 1.0,
            duration: 700,
          }
        ),

      ]).start(() => {
        this.setState({showLoading: true});
        resolve();
      });
    });
  }

  render() {
    const { showLoading } = this.state;
    const opacityVal = (showLoading) ? 1.0 : 0.0;
   
    return (
      <View style={[styles.topContainer, styles.container]}>
        <StatusBar barStyle="light-content" />
        <Animated.View style={[styles.logoWrapper, {width: this.state.logoWidth}]}>
          <Logo />
        </Animated.View>
        <ActivityIndicator size="small"
          color={'#000000'}
          style={[styles.loading, {opacity: opacityVal}]}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  topContainer: {
    backgroundColor: '#88ffff',
    height: '100%',
    width: '100%',
    flexDirection: 'column',
  },
  container: {
    justifyContent: 'center',
    alignItems: 'center',
  },
  logoWrapper: {
    flexDirection: 'row',
    justifyContent: 'flex-start',
    alignItems: 'center',
    height: 80,
    width: 80,
  },
  loading: {
    marginTop: 24,
    width: 20,
    height: 20,
  }
});

