import React, { useEffect } from 'react'

import Button from '@material-ui/core/Button'

import { useUser } from '../../contexts/UserCtx'
import { signOut, getUser } from '../../lib/auth'

export default function SignOutButton () {
  const { setUser } = useUser()

  useEffect(() => {
    const handleUser = async () => {
      const user = await getUser()
      setUser(user)
    }

    handleUser()
  }, [setUser])

  const handleSignOut = async () => {
    try {
      const res = await signOut()
      setUser(res)
    } catch (err) {
      console.error(err)
      setUser(null)
    }
  }

  return (
    <Button variant="contained" color="secondary" onClick={handleSignOut}>
      Sign out
    </Button>
  )
}
