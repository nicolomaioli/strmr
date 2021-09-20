import React, { useContext, useState, useMemo } from "react";

import PropTypes from "prop-types";

const UserCxt = React.createContext();

const useUser = () => {
  return useContext(UserCxt);
};

const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const providerValue = useMemo(() => ({ user, setUser }), [user, setUser]);

  return <UserCxt.Provider value={providerValue}>{children}</UserCxt.Provider>;
};

UserProvider.propTypes = {
  children: PropTypes.oneOfType([PropTypes.element, PropTypes.node]),
};

export { useUser, UserProvider };
