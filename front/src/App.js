import React, { Component } from 'react';
import './App.css';
import { withStyles } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';
import InputBase from '@material-ui/core/InputBase';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import SendButton from '@material-ui/icons/Send';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/List';
import ListItemText from '@material-ui/core/ListItemText';




const styles = theme =>({
  root: {

    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    flexWrap: 'wrap',
    margin: 'auto',
    backgroundColor: theme.palette.background.paper,

  },
  input: {
    marginLeft: 7,
    flex: 1,
  },
  SendButton: {
    padding: 10,
  },
  paper: {
    display: 'flex',
    alignItems: 'center',
    width: 400,
  },
});



class App extends Component {
  
  constructor(props){
    super(props)
    this.classes = props.classes
    this.state = {
      originalUrl : '',
      custom: '',
      shorten: ''
    }
    this.updateInput = this.updateInput.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleKeyPress = this.handleKeyPress.bind(this)
  }

  updateInput(event){
    this.setState({originalUrl : event.target.value})
  }
  
  handleSubmit(){
    if(this.state.originalUrl === "") {
      return;
    } else{
      console.log(JSON.stringify({
        Original_URL: this.state.originalUrl,
        custom: this.state.custom,
      }))

      fetch('/short/',{
        method: 'POST',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
        },
        body: JSON.stringify({
          Original_URL: this.state.originalUrl,
          custom: this.state.custom,
        })
      }).then(response => response.json())
      .then(data => data.data.short_url)
      .then(e => {
        if(e.length > 0){
          this.setState({shorten: e})
        }
      });
    }
    // alert(this.state.originalUrl)
  }

  handleKeyPress(e){
    if (e.keyCode === 13) {
      this.handleSubmit();
    }
  }

 

  render() {
    return (
      <div className={this.classes.root}>
        <Paper className={this.classes.paper} elevation={1}>
          <InputBase className={this.classes.input} onChange={this.updateInput} placeholder="Input Url" onKeyDown={this.handleKeyPress} />
          <IconButton className={this.classes.iconButton} aria-label="Send" onClick={this.handleSubmit}>
            <SendButton/>
          </IconButton>
          <Divider className={this.classes.divider} />
        </Paper>
        {this.state.shorten && (
          <div>
            <List component="nav">
              <ListItem href={this.state.shorten}>
                <ListItemText>{'http://localhost:9651/'+this.state.shorten}</ListItemText>
              </ListItem>
            </List>
          </div>
        )}
      </div>
    );
  }
}

export default withStyles(styles)(App)