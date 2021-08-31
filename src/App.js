import React, { useState, useEffect } from 'react'
import Amplify, { Auth } from 'aws-amplify'
import { AmplifySignOut, withAuthenticator } from '@aws-amplify/ui-react'
import awsconfig from './awsconfig'
import { Container, Typography } from '@material-ui/core'

Amplify.configure(awsconfig)

function App () {
  const [loggedIn, setLoggedIn] = useState(false)
  const [user, setUser] = useState({})

  const isLoggedIn = async () => {
    try {
      const user = await Auth.currentAuthenticatedUser()
      setUser(user)
      setLoggedIn(true)
    } catch (err) {
      setLoggedIn(false)
    }
  }

  useEffect(() => {
    isLoggedIn()
  }, [])

  return (
    <Container>
      <Typography variant="h1">
        Hello World! { loggedIn ? user.getUsername() : ''}
      </Typography>
      <AmplifySignOut />
    </Container>
  )
}

export default withAuthenticator(App)
