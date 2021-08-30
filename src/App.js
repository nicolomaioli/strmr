import React, { useState, useEffect } from 'react'
import Amplify, { Auth } from 'aws-amplify'
import { AmplifySignOut, withAuthenticator } from '@aws-amplify/ui-react'
import awsconfig from './aws-exports'

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
    <div>
      <h1>Hello World! { loggedIn ? user.getUsername() : ''}</h1>
      <AmplifySignOut />
    </div>
  )
}

export default withAuthenticator(App)
