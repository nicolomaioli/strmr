import React, { useEffect } from 'react'

import AppBar from '@material-ui/core/AppBar'
import { makeStyles } from '@material-ui/core/styles'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'

import { useUser } from '../contexts/UserCtx'
import { getUser } from '../lib/auth'
import SignInButton from './authentication/SignInButton'
import SignOutButton from './authentication/SignOutButton'

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
              : <SignInButton />
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
