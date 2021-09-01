import React, { useEffect } from 'react'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button'
import { makeStyles } from '@material-ui/core/styles'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { useUser } from '../contexts/UserCtx'
import SignOutButton from './authentication/SignOutButton'
import { signIn, getUser } from '../lib/auth'

const useStyles = makeStyles(() => ({
  title: {
    flexGrow: 1
  }
}))

export default function Nav () {
  const { user, setUser } = useUser()

  useEffect(() => {
    const handleUser = async () => {
      const user = await getUser()
      setUser(user)
    }

    handleUser()
  }, [setUser])

  const handleSignIn = async () => {
    try {
      const res = await signIn()
      setUser(res)
    } catch (err) {
      console.error(err)
      setUser(null)
    }
  }

  const classes = useStyles()

  return (
    <React.Fragment>
      <AppBar elevation={0}>
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            Strmr
          </Typography>
          {
            user
              ? <SignOutButton />
              : <Button variant="contained" color="primary" onClick={handleSignIn}>
                  Sign In
                </Button>
          }
        </Toolbar>
      </AppBar>
      {/*
        Would you believe me if I told you this is in the actual docs?
        https://material-ui.com/components/app-bar/#fixed-placement
      */}
      <Toolbar />
    </React.Fragment>
  )
}
