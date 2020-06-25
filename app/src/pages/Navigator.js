import React from 'react'
import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';

import Home from './home/Home';
import Login from './login/Login';

export default class Navigator extends React.Component{
    constructor() {
        super();
        this.state = {
            authenticated: false,
        };
    }

    render() {
	const { authenticated } = this.state;
	
        return(
            <BrowserRouter>
                <Switch>
		<Route exact path="/" render={() => (
			authenticated ? (
				<Redirect to="/home" />
			) : (
				<Redirect to="/login" />
			)
		)}/>
                    <Route exact path="/home" component={Home}/>
                    <Route exact path="/login" component={Login}/>
                </Switch>
            </BrowserRouter>
        );
    }
}
