import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Home from './containers/Home';
import Info from './containers/Info';
import Settings from './containers/Settings';

// Wrapper
import WrapMainContent from './components/WrapMainContent'


// 不明なルート
const NotFound = () => {
  return(
    <h2>Not Found Page</h2>
  );
}

class App extends React.Component {
  render(){
    return(
      <div className="App">
        <Switch>
          <Route exact path="/" component={WrapMainContent(Home)} />
          <Route exact path="/info" component={WrapMainContent(Info)} />
          <Route exact path="setting" component={WrapMainContent(Settings)} />
          <Routte component={WrapMainContent(NotFound)} />
        </Switch>
      </div>
    );
  };
};

export default App;
