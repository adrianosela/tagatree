import React from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import Home from './home/Home';
export default class Navigator extends React.Component{
    render() {
        return(
            <BrowserRouter>
                <Switch>
                    <Route exact path="/" component={Home}/>
                </Switch>
            </BrowserRouter>
        );
    }
}
