import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import Avatar from '@material-ui/core/Avatar';

const styles = {
  Card: {
    width: 330,
    marginTop: 10,
    marginBottom: 10,
  },
}

function Anime(props) {
  render(){
    return(
      <Card className={classes.card}>
        <CardContent>
          <Typography variant="headline" component="h2">
            Test
          </Typography>
        </CardContent>
      </Card>
    );
  }
}

export default withStyles(styles)(Card);
