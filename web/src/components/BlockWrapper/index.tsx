import React from 'react';
import styled from 'styled-components';

export interface BlockWrapperProps {
  children: React.ReactNode
}

export const BlockWrapper: React.FC<BlockWrapperProps> = ({
  children,
  ...props
}) => {
  return (
    <BlockWrapperStyled {...props}>{children}</BlockWrapperStyled>
  )
}

const BlockWrapperStyled = styled.div`
  display: flex;
  justify-content: space-between;
  width: 370px;   
`