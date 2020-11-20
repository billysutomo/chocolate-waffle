import React from "react";
import styled from "styled-components";
import { Element, Types } from "../../components/Element";
import BasicLayout from "../../components/BasicLayout";

const ContainerStyled = styled.div`
  max-width: 400px;
  margin: 0 auto;
  padding-bottom: 16px;
  padding-top: 16px;
`;

const SvgStyled = styled.svg`
  position: relative;
  display: block;
  font-size: 2.7em;
  width: 1em;
  height: 100%;
  line-height: 100%;
  margin: 0 auto;
  margin-top: -0.1rem;
  color: #ffffff;
`;

const SvgContainerStyled = styled.div`
  border-color: rgba(255, 255, 255, 0.5);
  border: 2px dashed #ffffff;
  border-radius: 50px;
  width: 100px;
  height: 100px;
  margin: 0 auto;
  margin-bottom: 15px;
`;

const TitleStyled = styled.div`
  width: 100%;
  text-align: center;
  h1 {
    border-bottom: 2px dashed #fffdfd;
    display: inline-block;
    font-size: 28px;
    margin: 0;
    padding: 0;
    margin-bottom: 16px;
    text-align: center;
    color: white;
  }
`;
/* icon list
- whatsapp
- facebook
- telegram
- skype
- viber
- email
- phone
*/

const dataDummy = [
  {
    type: "messenger",
    messenger_type: "whatsapp",
    value: "081313131313",
    text: "text gua lho",
  },
  {
    type: "messenger",
    messenger_type: "facebook",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "telegram",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "skype",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "viber",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "email",
    value: "billysutomo.53@gmail.com",
    subject: "subject email",
  },
  {
    type: "messenger",
    messenger_type: "phone",
    value: "081313131313",
  },
];

const Component: React.FC = () => {
  const renderMessenger = () => {
    return dataDummy.map((a) => {
      let messenger: Types;
      if (a.messenger_type === "messenger") {
        messenger = Types.messenger;
      } else if (a.messenger_type === "block") {
        messenger = Types.block;
      } else if (a.messenger_type === "social_link") {
        messenger = Types.social_link;
      } else {
        messenger = Types.nothing;
      }
      return (
        <Element active elementType={messenger}>
          {a.value}
        </Element>
      );
    });
  };

  return (
    <BasicLayout backgroundColor="#cf8383">
      <ContainerStyled>
        <SvgContainerStyled>
          <SvgStyled>
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fal"
              data-icon="camera"
              className="svg-inline--fa fa-w-16 fa-camera"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
            >
              <path
                fill="currentColor"
                d="M324.3 64c3.3 0 6.3 2.1 7.5 5.2l22.1 58.8H464c8.8 0 16 7.2 16 16v288c0 8.8-7.2 16-16 16H48c-8.8 0-16-7.2-16-16V144c0-8.8 7.2-16 16-16h110.2l20.1-53.6c2.3-6.2 8.3-10.4 15-10.4h131m0-32h-131c-20 0-37.9 12.4-44.9 31.1L136 96H48c-26.5 0-48 21.5-48 48v288c0 26.5 21.5 48 48 48h416c26.5 0 48-21.5 48-48V144c0-26.5-21.5-48-48-48h-88l-14.3-38c-5.8-15.7-20.7-26-37.4-26zM256 408c-66.2 0-120-53.8-120-120s53.8-120 120-120 120 53.8 120 120-53.8 120-120 120zm0-208c-48.5 0-88 39.5-88 88s39.5 88 88 88 88-39.5 88-88-39.5-88-88-88z"
              />
            </svg>
          </SvgStyled>
        </SvgContainerStyled>
        <TitleStyled>
          <h1>Title here</h1>
        </TitleStyled>
        {/* <Element active={false}>+ Add Element</Element>
        <Element active={false}>+ Add Element</Element>
        <Element active={false}>+ Add Element</Element> */}
        {renderMessenger()}
      </ContainerStyled>
    </BasicLayout>
  );
};

export default Component;
