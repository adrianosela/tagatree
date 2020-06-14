import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

import Splash from './src/pages/splash/Splash';
import Login from './src/pages/login/Login';
import Home from './src/pages/home/Home';
import Tag from './src/pages/tag/Tag';

const Stack = createStackNavigator();

function App() {
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

export default App;
