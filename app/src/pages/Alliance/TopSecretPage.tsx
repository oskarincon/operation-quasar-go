
import { LayoutContainer } from '@/styled-components';
import { TopSecret } from '@/components';

export const TopSecretPage = () => {

  return (
    <>
    <LayoutContainer>
      <div style={{ fontSize: '2rem'}}>
        <h4>Top Secret</h4>
      </div>
    </LayoutContainer>
    <div style={{ fontSize: '2rem' }}>
      <TopSecret data={[{index: 0, name: "kenobi", text: "", cord: ""}, {index: 1, name: "skywalker", text: "", cord: ""}, {index: 2, name: "sato", text: "", cord: ""}]}></TopSecret>
    </div>
  </>
  );
};

export default TopSecretPage;
