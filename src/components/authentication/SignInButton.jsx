import React, { useEffect } from 'react'

import Button from '@material-ui/core/Button'

import { useUser } from '../../contexts/UserCtx'
import { signIn, getUser } from '../../lib/auth'

export default function SignInButton () {
  const { setUser } = useUser()

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

  return (
    <Button variant="contained" color="primary" onClick={handleSignIn}>
      Sign In
    </Button>
  )
}
