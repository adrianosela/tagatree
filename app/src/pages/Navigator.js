import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import Splash from './splash/Splash';
import Login from './login/Login';
import Home from './home/Home';
import Tag from './tag/Tag';

const Stack = createStackNavigator();

export default class Navigator extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <NavigationContainer>
        <Stack.Navigator initialRouteName="Splash">
          <Stack.Screen options={{headerShown: false}} name='Splash' component={Splash} />
          <Stack.Screen options={{headerShown: false}} name='Login' component={Login} />
          <Stack.Screen options={{headerShown: false}} name='Home' component={Home} />
          <Stack.Screen name='Tag' component={Tag} />
        </Stack.Navigator>
      </NavigationContainer>
    );
  }
}
