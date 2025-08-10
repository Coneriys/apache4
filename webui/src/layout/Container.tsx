import { Flex, styled } from '@apache4labs/faency'

import breakpoints from 'utils/breakpoints'

export default styled(Flex, {
  flexGrow: 1,
  margin: '0 24px',

  [`@media (min-width: ${breakpoints.laptopL})`]: {
    maxWidth: `calc(${breakpoints.laptopL} - 96px)`,
    margin: '0 auto',
  },
})
