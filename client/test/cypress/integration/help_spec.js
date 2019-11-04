// DONE
describe('Help page - check the links and picture', function () {
  it('Feature link works', function () {
    cy.visit('/help')
    cy.get('a').contains('Describe your need').should('have.attr', 'href', 'https://docs.google.com/forms/d/e/1FAIpQLSeh5Gw8iH9DlLQN3tjF66rrDLn5dkPnbi3V6vSjKAu0QPXnag/viewform?usp=sf_link')
  })

  it('Bug link works', function () {
    cy.visit('/help')
    cy.get('a').contains('Report it').should('have.attr', 'href', 'https://docs.google.com/forms/d/e/1FAIpQLSfExGW2_D3YI4SjQ3-7SC7gys-b1AwywY-TBmITlDX644-Upg/viewform?usp=sf_link')
  })

  it("Help image is visible", () => {
    cy.get('div[role=img]').should('be.visible')
  })
})
