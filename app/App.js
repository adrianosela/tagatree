import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

import Splash from './components/pages/splash/Splash';
import Login from './components/pages/login/Login';
import Home from './components/pages/home/Home';
import Tag from './components/pages/tag/Tag';

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
