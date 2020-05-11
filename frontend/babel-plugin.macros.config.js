const isDev = process.env.NODE_ENV !== 'production'

export const styledComponents = {
  fileName: isDev,
  displayName: isDev,
}
