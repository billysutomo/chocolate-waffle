import React from 'react';
import styled, { css } from 'styled-components';

export interface BlockProps {
  children: string;
  active: boolean;
  backgroundColor?: string;
  size?: 'small' | 'medium' | 'large';
}

export const Block: React.FC<BlockProps> = ({
  children,
  active,
  backgroundColor,
  size = 'large',
  ...props
}) => {
  return (
    <BlockStyled
      size={size}
      active={active}
      backgroundColor={backgroundColor}
      {...props}
    >{children}</BlockStyled>
  )
}

const BlockStyled = styled.button<BlockProps>`
  border: 2px dashed #607D8B;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
  background-color: ${p => p.backgroundColor};
  :hover{
    cursor: pointer;
  }
  :active{
    background: #2ab8b96b;
  }
  :focus{
    outline: 0;
  }
  ${p => p.size === 'medium' && css`
    width: 177px;
  `}
  ${p => p.size === 'small' && css`
    width: 112.66px;
  `}
  ${p => p.active && css`
    border: none;
  `}
`