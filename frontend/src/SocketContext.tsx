import React, { createContext, useContext, useState } from "react"

type SocketContextType = {
  socket: WebSocket | null
  setSocket: (socket: WebSocket) => void
}

const SocketContext = createContext<SocketContextType>({
  socket: null,
  setSocket: () => {},
})

export const SocketProvider: React.FC = ({ children }) => {
  const [socket, setSocket] = useState<SocketContextType["socket"]>(null)

  return (
    <SocketContext.Provider value={{ socket, setSocket }}>
      {children}
    </SocketContext.Provider>
  )
}

export const useSocket = () => useContext(SocketContext)
