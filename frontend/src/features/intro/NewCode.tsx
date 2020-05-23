import {
  CircularProgress,
  Typography,
  TypographyProps,
} from "@material-ui/core"
import React, { useEffect, useState } from "react"
import styled from "styled-components/macro"
import { useSocket } from "SocketContext"

const _CodeDigit: React.FC<TypographyProps> = (props) => (
  <Typography component="span" {...props} />
)
const StyledCodeDigit = styled(_CodeDigit)`
  font-size: 4.5rem;
  margin: ${(p) => p.theme.spacing(0, 1)};
`
const StyledCircularProgress = styled(CircularProgress)`
  margin: ${(p) => p.theme.spacing(2)}px;
`

const NewCode = () => {
  const [code, setCode] = useState<string | null>(null)

  const { setSocket } = useSocket()

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:3000/code/new")

    socket.onmessage = function (event) {
      setCode(event.data)
    }

    setSocket(socket)
  }, [setSocket])

  return (
    <>
      {code && (
        <>
          {code.split("").map((digit, i) => (
            <StyledCodeDigit key={i}>{digit}</StyledCodeDigit>
          ))}
          <Typography>Enter this code on another device...</Typography>
        </>
      )}
      <StyledCircularProgress disableShrink={!!code} />
    </>
  )
}

export default NewCode
