import React from 'react'
import PropTypes from 'prop-types'
import { Box, CircularProgress, Typography } from '@material-ui/core'

const StatusHandler = (props) => {
    const {
        children,
        message,
        status,
        height,
        width,
        style,
        className,
        ...other
    } = props
    return (
        <Box height={height} width={width} style={style} {...other}>
            {status === 'fail' && (
                <Box
                    display="flex"
                    alignItems="center"
                    justifyContent="center"
                    height="100%"
                    width="100%"
                >
                    <Typography variant="h6">{message}</Typography>
                </Box>
            )}
            {status === 'initial' && (
                <Box
                    className={`status-handler-message`}
                    display="flex"
                    alignItems="center"
                    justifyContent="center"
                    height="100%"
                    width="100%"
                >
                    <div className="logo" align="center">
                        <div align="center">
                            <Typography variant="h6">{message}</Typography>
                        </div>
                    </div>
                </Box>
            )}
            {status === 'loading' && (
                <Box
                    className="status-handler-loading"
                    display="flex"
                    alignItems="center"
                    justifyContent="center"
                    flexDirection="column"
                    height="100%"
                    width="100%"
                    minHeight="inherit"
                >
                    <CircularProgress />
                    <Typography variant="h6">{message}</Typography>
                </Box>
            )}
            {status === 'success' && children}
        </Box>
    )
}

StatusHandler.defaultProps = {
    height: '100%',
    width: '100%',
    minHeight: '150px',
    message: 'Please click on the table above to display details.',
    status: 'initial',
    // initialRenderChild: false,
}

StatusHandler.propTypes = {
    className: PropTypes.string,
    height: PropTypes.string,
    width: PropTypes.string,
    style: PropTypes.object,
    children: PropTypes.node,
    message: PropTypes.string,
    status: PropTypes.oneOf(['initial', 'success', 'fail', 'loading']),
}

export default StatusHandler
