import React from "react";
import { withStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import Icon from "@material-ui/core/Icon";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Api from "./api/Api";

const styles = theme => ({
  root: {
    flexGrow: 1,
    padding: "5em"
  },
  paper: {
    padding: theme.spacing(2),
    margin: "auto",
    maxWidth: 500
  },
  image: {
    width: 128,
    height: 128
  },
  img: {
    margin: "auto",
    display: "block",
    maxWidth: "100%",
    maxHeight: "100%"
  },
  heroContent: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(8, 0, 6)
  },
  heroButtons: {
    marginTop: theme.spacing(4)
  }
});

class App extends React.Component {
  sendNow() {
    Api.tasks.sendNow();
  }
  scheduleTask() {
    Api.tasks.schedule();
  }
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <div className={classes.heroContent}>
          <Container maxWidth="sm">
            <Typography
              component="h1"
              variant="h2"
              align="center"
              color="textPrimary"
              gutterBottom
            >
              Task Scheduler
            </Typography>
            <Typography
              variant="h5"
              align="center"
              color="textSecondary"
              paragraph
            >
              Run Now or Schedule Task
            </Typography>
          </Container>
        </div>
        <Grid container spacing={3}>
          <Grid item xs={6}>
            <Button
              variant="contained"
              color="primary"
              className={classes.button}
              fullWidth
              onClick={this.sendNow.bind(this)}
              endIcon={<Icon>send</Icon>}
            >
              Send
            </Button>
          </Grid>
          <Grid item xs={6}>
            <Button
              variant="contained"
              color="primary"
              fullWidth
              className={classes.button}
              endIcon={<Icon>access_time</Icon>}
              onClick={this.scheduleTask.bind(this)}
            >
              Schedule
            </Button>
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default withStyles(styles)(App);

/*
<div className="App">
        <header className="App-header">
        </header>
      </div>
*/
