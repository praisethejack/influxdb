// Libraries
import React, {PureComponent} from 'react'
import {withRouter, RouteComponentProps} from 'react-router-dom'

// Components
import {Overlay} from '@influxdata/clockface'
import VariableFormContext from 'src/variables/components/VariableFormContext'
import GetResources from 'src/resources/components/GetResources'

// Types
import {ResourceType} from 'src/types'

type Props = RouteComponentProps<{orgID: string}>

class CreateVariableOverlay extends PureComponent<Props> {
  public render() {
    return (
      <Overlay visible={true}>
        <Overlay.Container maxWidth={1000}>
          <Overlay.Header
            title="Create Variable"
            onDismiss={this.handleHideOverlay}
          />
          <Overlay.Body>
            <GetResources resources={[ResourceType.Variables]}>
              <VariableFormContext onHideOverlay={this.handleHideOverlay} />
            </GetResources>
          </Overlay.Body>
        </Overlay.Container>
      </Overlay>
    )
  }

  private handleHideOverlay = () => {
    const {
      history,
      match: {
        params: {orgID},
      },
    } = this.props

    history.push(`/orgs/${orgID}/settings/variables`)
  }
}

export {CreateVariableOverlay}
export default withRouter(CreateVariableOverlay)
